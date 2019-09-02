<template lang="pug">
  b-navbar.top-nav(type="dark" variant="dark")
    //- a.navbar-brand
    //-   img(src="logo.svg" width="30" height="30" alt="")
    b-navbar-brand
      | maccabee.io
    b-navbar-toggle(target="nav-collapse")

    b-collapse#nav-collapse(is-nav)
      b-navbar-nav
        b-nav-item(v-for="route in routes" :key="route.path" :to="route.path" exact-active-class="active")
          | {{route.name}}
</template>

<style>
/* .top-nav {
  background-color: #563d7c;
  min-height: 3rem;
} */
</style>

<script>
  export default {
    name: "navbar",
    data: function() {
      return {
        routes: this.$router.options.routes
      }
    }
  }
</script>
