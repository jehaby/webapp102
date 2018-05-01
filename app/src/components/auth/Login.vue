<template>
<div class="column is-half">
  <form>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Name</label>
      </div>
      <div class="field-body">
        <div class="field is-narrow">
          <div class="control">
            <input v-model="user.name" class="input" type="text">
          </div>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Password</label>
      </div>
      <div class="field-body">
        <div class="field is-narrow">
          <div class="control">
            <input v-model="user.password" class="input" type="password">
          </div>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label">
        <!-- Left empty for spacing -->
      </div>
      <div class="field-body">
        <div class="field">
          <div class="control">
            <button v-on:click="login" class="button is-primary">
              Login
            </button>
          </div>
        </div>
    </div>
  </div>
    
  </form>
</div>
</template>

<script>
import { loginRequest } from '../../api/auth.js'

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
  methods: {
    async login () {
      // TODO: form validation
      let resp = {}
      try {
        resp = await loginRequest({ ...this.user })
        this.$store.commit('user', resp.user)
        this.$store.commit('jwtToken', resp.token)
        this.$router.push('/profile')
      } catch (e) {
        console.log('foooo', e)
        // TODO: better errors
        return this.$store.dispatch('error', 'Login failed')
      }
    }
  }
}
</script>

<style scoped>

</style>
