<template lang="pug">
  b-navbar(type="dark" variant="dark")
    b-navbar-brand
      | exofish

    b-navbar-toggle(target="nav-collapse")

    b-collapse#nav-collapse(is-nav)
      b-navbar-nav
        b-nav-item(v-for="route in routes" :key="route.path" :to="route.path" exact-active-class="active")
          | {{route.name}}
</template>

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
