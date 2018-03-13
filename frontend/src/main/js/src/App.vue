<template>
  <div id="app">
    <header>
      <div class="header-wrapper">
        <img src="http://acmelogos.com/images/logo-8.svg" alt="" class="logo">
        <!--<nav>
          <router-link to="/blog" active-class="grey">+ Feed</router-link>
        </nav>-->
      </div>
    </header>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  data() {
    return {
      backend:'http://127.0.0.1:8000'
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 10px;
}
</style>
