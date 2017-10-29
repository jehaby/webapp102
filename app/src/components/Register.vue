<template>
<div>
  <p> {{ count }}</p>
  <h5>register</h5>
  <form>
    <p>
      name
      <input type="text"/>
    </p>
    <p>
      email
      <input type="email"/>
    </p>
    <p>
      password
      <input type="password"/>
    </p>
    <p>
      <input type="button" value="ok"/>
    </p>
  </form>
</div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: 'Register',
    computed: mapState(['count'])
  }
</script>

<style scoped>

</style>
