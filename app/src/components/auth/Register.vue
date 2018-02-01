<template>
<div class="column is-half">
  <form>

    <div class="field">
      <label class="label">Name</label>
      <div class="control">
        <input v-model="user.name" class="input" type="text">
      </div>
    </div>

    <div class="field">
      <label class="label">Email</label>
      <div class="control">
        <input v-model="user.email" class="input" type="email">
      </div>
    </div>

    <div class="field">
      <label class="label">Password</label>
      <div class="control">
        <input v-model="user.password" class="input" type="password">
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button v-on:click="register" class="button is-link">Register</button>
      </div>
    </div>

  </form>
</div>
</template>

<script>
  import { registerRequest } from '../../api/auth.js'

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
