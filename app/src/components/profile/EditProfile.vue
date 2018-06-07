<template>
<div>

<form>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">email</label>
      </div>
      <div class="field-body">
        <div class="field is-narrow">
          <div class="control">
            <input v-model="user.email" class="input" type="email">
          </div>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">phones</label>
      </div>
      <div class="field-body">
        <div class="field is-narrow">
          <div class="control">
            <input v-model="user.phone" class="input" type="phone">
          </div>
        </div>
      </div>
    </div>
    

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">password</label>
      </div>
      <div class="field-body">
        <div class="field is-narrow">
          <div class="control">
            <button v-on:click="changePassword" class="button">
              change_password
            </button>
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
import { userGetRequest, userUpdateRequest } from '@/api/user.js'
export default {
  name: 'EditProfile',
  data () {
    return {
      user: {
        name: ''
      }
    }
  },
  methods: {
    async getUser () {
      try {
        const resp = await userGetRequest(this.$store.state.auth.auth.user.uuid)
        this.user = resp.data
      } catch (e) {
        console.log(e)
      }
    },
    async update () {
      try {
        const resp = await userUpdateRequest()
        console.log(resp)
      } catch (e) {
        console.log(e)
      }
    },
    changePassword () {
      this.$router.push()
    },
    send () {

    }
  },
  created () {
    console.log('in created')
    this.getUser()
  }
}
</script>

