<template>
<div>
  <p> {{ count }}</p>
  <h5>register</h5>
  <form>
    <p>
      name
      <input v-model="user.name" type="text"/>
    </p>
    <p>
      email
      <input v-model="user.email" type="email"/>
    </p>
    <p>
      password
      <input v-model="user.password" type="password"/>
    </p>
    <p>
      <input type="button" v-on:click="register" value="ok"/>
    </p>
  </form>
</div>
</template>

<script>
  import { mapState } from 'vuex'
  import { register } from '../api'

  export default {
    name: 'Register',
    data () {
      return {
        user: {
          name: '',
          email: '',
          password: ''
        }
      }
    },
    computed: mapState(['count']),
    methods: {
      register () {
        register({...this.user})
      }
    }
  }
</script>

<style scoped>

</style>
