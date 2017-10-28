<template>
  <div>
    <h5> login </h5>
    <form>
      <p>
        name
        <input type="text"/>
      </p>
      <p>
        password
        <input type="password"/>
      </p>
      <p>
        <input type="button" value="login"/>
      </p>
    </form>
  </div>
</template>

<script>
  export default {
    name: 'Login',
    data () {
      return {
        msg: '?'
      }
    }
  }
</script>

<style scoped>

</style>
