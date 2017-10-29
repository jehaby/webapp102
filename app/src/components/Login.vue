<template>
  <div>
    <p> {{ count }} </p>
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
  import { mapState } from 'vuex'

  export default {
    name: 'Login',
    data () {
      return {
        msg: '?'
      }
    },
    computed: mapState({
      // or state => state.count
      count: 'count'
    })
  }
</script>

<style scoped>

</style>
