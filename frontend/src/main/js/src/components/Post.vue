<template>
  <div class="post" v-if="post">
    <img :src="post.image" alt="">
    <h2>{{ post.title }}</h2>
    <p>{{ post.body }}</p>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    data() {
      return {
        post: null
      }
    },
    created() {
      this.postId = this.$route.params.id
    },
    mounted () {
      axios.get("http://127.0.0.1:8000/news/" + this.postId)
      .then(response => {
        this.post = response.data;
        console.log(response)
      })
      .catch(e => {
        console.log('error', e)
      })
    }
  }
</script>
