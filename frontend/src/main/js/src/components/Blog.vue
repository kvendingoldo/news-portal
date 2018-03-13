<template>
  <div class="blog">
    <div class="blog__post" v-for="(post, index) in this.results">
      <img :src="post.image" alt="">
      <h3>{{ post.title }}</h3>
      <router-link :to="{ name: 'post', params: { id: post.id } }">Подробнее</router-link>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    data() {
      return {
        results: this.getPosts()
      }
    },
    methods: {
      getPosts: function() {
        axios.get('http://127.0.0.1:8000/news')
        .then((response) => {
          this.results = response.data;
          console.log(response)
        }).catch((error) => {
          console.log(error);
        });
      }
    }
  }
</script>
