<template lang="pug">
  div.home
    h2 You've made it!
    br
    h5 Click #[b-link(to="/shuttles") here] for the shuttles page.
    br
    h5 Click #[b-link(to="/about") here] to read some stuff about me.
</template>

<script>
export default {
  name: 'Home'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
a {
  color: #42b983;
}
</style>
