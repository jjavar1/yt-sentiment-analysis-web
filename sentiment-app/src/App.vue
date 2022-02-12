<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-12 p-0">
        <div class="jumbotron min-vh-100 text-center m-0 d-flex flex-column">
          <h1 class="display-1" style="margin-top: 10px">Sentiment</h1>
          <form v-on:submit.prevent="generateSentiment().then(getSentiment)">
            <p style="margin-top: 30px">
              <input
                :disabled="isDisabled"
                v-model="video_ID"
                type="text"
                class="form-control input-lg mx-auto"
                id="youtube-url"
                placeholder="youtube addess"
                style="width: 15%"
              />
            </p>
            <p style="margin-top: 30px">
              <button class="btn btn-primary btn-lg" :disabled="isDisabled" v-on:click="isHidden = !isHidden">get sentiment</button>
            </p>
            <p style="margin-top: 28px"><span>{{ sentimentComment }}</span></p>
            <p style="margin-top: 5px" v-if="!isHidden">Refresh to generate new sentiment</p>
          </form>
          <p style="margin-top: 10px"></p>
           <div class="btn-groupd">
                <button class="btn btn-secondary mr-1" role="button" v-if="!isHidden" @click="drawChart">Total Sentiment</button>
                <button class="btn btn-secondary mr-1" role="button" v-if="!isHidden">ML Sentiment</button>
                <button class="btn btn-secondary mr-1" role="button" v-if="!isHidden">Lexicon Sentiment</button>
          </div>
          <canvas v-if="!isHidden" id = 'example' height = '90' width = '400' ></canvas>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import axios from "axios";
import { defineComponent } from "vue";
import { Chart, registerables } from "chart.js";

export default defineComponent({
  name: "app",
  data() {
    return {
      commentSentiment: "",
      video_ID: "",
      sentimentRank: 0,
      sentimentComment: "",
      totalPositiveSent: 0,
      totalNegativeSent: 0,
      totalPosMLSen: 0,
      totalNegMLSen: 0,
      isHidden: true,
      isDisabled: false
    };
  },
  methods: {
    generateSentiment: function() {
      this.isDisabled = true
      return axios
        .post("http://localhost:3000/api/yt", { video_ID: `${this.video_ID}` })
        .then((response) => {
          this.video_ID = response.data.video_ID;
        })
        .catch((error) => {
          window.alert(`The api returned an error: ${error}`);
        });
        
    },
    getSentiment: function() {
      return axios
        .get("http://localhost:3000/api/yt/get")
        .then((response) => {
          this.sentimentRank = response.data.sentScore.SentScore;
          this.totalPositiveSent =
            response.data.sentScore.TotalPositiveComments;
          this.totalNegativeSent =
            response.data.sentScore.TotalNegativeComments;
          this.totalPosMLSen = response.data.sentScore.
          console.log(this.totalPositiveSent)
          if (this.sentimentRank == 0) {
            return;
          } else if (this.sentimentRank == 1) {
            return (this.sentimentComment = `This video has a good sentiment`);
          } else if (this.sentimentRank == -1) {
            return (this.sentimentComment = `This video has a bad sentiment`);
          }
          console.log(this.totalPositiveSent)
          
        })
   
        .catch((error) => {
          window.alert(`The api returned an error: ${error}`);
        });
    },
    drawChart: function() {
      console.log("genchart" + this.totalPositiveSent)
      Chart.register(...registerables)
      let canvas: any = document.getElementById('example')
        let ctx = canvas.getContext('2d')
        const chrt = new Chart(ctx, {
        type: 'bar',
        data: {
        labels: ["Positive Sentiment", "Negative Sentiment"],
        datasets: [
          {
            label: "Overall Sentiment",
            backgroundColor: ["green", "red"],
            data: [this.totalPositiveSent, this.totalNegativeSent],
          },
          ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: true,
      }
      });
        chrt.render();
    },
    drawPieChart: function() {
      Chart.register(...registerables)
      let canvas: any = document.getElementById('example')
      let ctx = canvas.getContext('2d')
      const chrt = new Chart(ctx, {
        type: 'pie',
        data: {
        labels: ["Positive Sentiment", "Negative Sentiment"],
        datasets: [
          {
            label: "Overall Sentiment",
            backgroundColor: ["green", "red"],
            data: [this.totalPositiveSent, this.totalNegativeSent],
          },
          ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: true,
      }
      });
        chrt.render();
    },
  },
});
</script>
