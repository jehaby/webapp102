<template>
  <div>
    <p> {{ count }} </p>
    <h5> login </h5>
    <form>
      <p>
        name
        <input v-model="user.name" type="text"/>
      </p>
      <p>
        password
        <input v-model="user.password" type="password"/>
      </p>
      <p>
        <input type="button" v-on:click="login" value="login"/>
      </p>
    </form>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import { loginRequest } from './../api'

  export default {
    name: 'Login',
    data () {
      return {
        user: {
          name: '',
          password: ''
        }
      }
    },
    computed: mapState({
      // or state => state.count
      count: 'count'
    }),
    methods: {
      login () {
        loginRequest({...this.user})
      }
    }
  }
</script>

<style scoped>

</style>
