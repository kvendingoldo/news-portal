<template>
  <div class="container" id="app">

    <section class="callout secondary">
      <h5 class="text-center">Filter by Category</h5>
      <form>
        <div class="row">
          <div class="large-6 columns">
            <select v-model="this.section">
              <option v-for="section in this.sections" :value="section">{{ section }}</option>
            </select>
          </div>
          <div class="medium-6 columns">
            <a @click="getPosts(section)" class="button expanded">Retrieve</a>
          </div>
        </div>
      </form>
    </section>

    <h5 class="text-center uppercase" v-if="!this.$parent.loading">{{ this.$parent.title }}</b></h5>
    <div v-if="this.loading" class="loader">
      <img src="http://cdnjs.cloudflare.com/ajax/libs/semantic-ui/0.16.1/images/loader-large.gif" alt="loader">
    </div>

    <NewsList v-if="!this.loading" :results="this.$parent.results"></NewsList>

    <Blog></Blog>
  </div>
</template>


<script>
import Blog from './Blog'

export default{
    data() {
      return {
        section: 'nature',
        sections: ['nature', 'misc'],
        loading: false
      }
    },
    components:{
         'Blog': Blog
    }
}
</script

<style>
		.container {
			padding: 20px 0
		}
		.column:last-child:not(:first-child), .columns:last-child:not(:first-child) {
			float: left;
		}
		.main_heading {
			margin: 10px 0 30px;
		}
		.callout {
			margin-bottom: 40px;
		}
		select {
			text-transform:capitalize;
		}
		.loader {
			text-align: center;
			padding: 40px;
		}
		.uppercase {
			text-transform: uppercase;
		}
</style>
