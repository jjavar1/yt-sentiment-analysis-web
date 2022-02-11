<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-12 p-0">
        <div class="jumbotron min-vh-100 text-center m-0 d-flex flex-column">
          <h1 class="display-1" style="margin-top: 130px">Sentiment</h1>
          <form v-on:submit.prevent="generateSentiment().then(getSentiment())">
          <p style="margin-top: 50px">
            <input
              v-model="video_ID"
              type="text"
              class="form-control input-lg mx-auto"
              id="youtube-url"
              placeholder="youtube addess"
              style="width: 15%"
            />
          </p>
          <p style="margin-top: 50px">
            <button class="btn btn-primary btn-lg"
              >get sentiment</button
            >
          </p>
          <span>{{ sentimentRank }}</span>
          </form>
          
        </div>
      </div>
    </div>
  </div>
</template>
<script lang = "ts">
import axios from 'axios'
import { defineComponent } from "vue";

export default defineComponent({
  name: "app",

  data() {
    return {
      commentSentiment: '',
      video_ID: '',
      sentimentRank: ''
    };
  },
  methods: {
    generateSentiment() { 
      return axios.post("http://localhost:3000/api/yt", { video_ID: `${this.video_ID}` }).then((response) => {
       this.video_ID = response.data.video_ID;
       console.log(response)
    }).catch((error) => {
      window.alert(`The api returned an error: ${error}`)
    })
  },
    getSentiment() {
      axios.get("http://localhost:3000/api/yt/get").then((response) => {
        this.sentimentRank = response.data
        console.log(this.sentimentRank)
      }).catch((error) => {
      window.alert(`The api returned an error: ${error}`)
    })
    }
  }
})
</script>
