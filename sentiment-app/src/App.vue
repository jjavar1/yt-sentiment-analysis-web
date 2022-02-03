<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-12 p-0">
        <div class="jumbotron min-vh-100 text-center m-0 d-flex flex-column">
          <h1 class="display-1" style="margin-top: 130px">Sentiment</h1>
          <form v-on:submit.prevent="generateSentiment">
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
          <span>{{ commentSentiment }}</span>
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
      video_ID: ''
    };
  },
  methods: {
    generateSentiment() {
      
      axios.post("http://localhost:3000/api/yt", { video_ID: `${this.video_ID}` }).then((response) => {
       console.log(response.data)
    }).catch((error) => {
      window.alert(`The api returned an error: ${error}`)
    })
}
  }
})
</script>
