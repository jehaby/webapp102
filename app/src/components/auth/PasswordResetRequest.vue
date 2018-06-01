<template>

<div class="column">
  <form>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">name_or_email</label>
      </div>
      <div class="field-body">
        <div class="field is-narrow">
          <div class="control">
            <input v-model="name" class="input" type="text">
          </div>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <!-- Left empty for spacing -->
      </div>
      <div class="field-body">
        <div class="field">
          <div class="control">
            <button v-on:click="send" class="button is-primary">
              send
            </button>
          </div>
        </div>
    </div>

  </div>

  </form>
</div>

</template>

<script>
import { passwordResetRequest } from '@/api/auth.js'

export default {
  name: 'PasswordResestRequest',
  data () {
    return {
      name: ''
    }
  },
  methods: {
    async send () {
      try {
        await passwordResetRequest({ name: this.name })
        this.$store.dispatch('showMsg', {type: 'info', text: 'password_reset_request_success'})
        this.$router.push('/')
      } catch (e) {
        this.$store.dispatch('error', 'failed.try_again_later')
      }
    }
  }
}
</script>

