<template>
<div>
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
  import { registerRequest } from '../api/auth.js'

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
    methods: {
      async register () {
        // TODO: form validation
        let resp = {}
        try {
          resp = await registerRequest({...this.user})
        } catch (err) {
          this.$store.dispatch('error', 'Registration failed!')
          return
        }
        this.$store.commit('user', resp.user)
        this.$store.commit('jwtToken', resp.token)
        this.$router.push('/profile')
      }
    }
  }
</script>

<style scoped>

</style>
