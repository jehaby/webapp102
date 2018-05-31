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
      try {
        await this.$store.dispatch('login', { ...this.user })
        // TODO: show success message
        this.$router.push('/profile')
      } catch (e) {
        console.log('error logging in', e)
        // TODO: better errors
        this.$store.dispatch('error', 'Login failed')
      }
    }
  }
}
</script>

<style scoped>

</style>
