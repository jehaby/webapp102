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
        <div class="field">
          <div v-for="phone in user.phones" class="field has-addons">
            <p class="control">
              <a class="button is-static ">+{{ phone.country_code + phone.number}}</a>
            </p>

            <p class="control">
              <a v-on:click="removePhone(phone.uuid)" class="button"> X </a>
            </p>
          </div>
          <div class="field is-narrow">
            <div class="field has-addons">
              <template v-if="phoneCreation">
                <!-- TODO: phone input -->
                <p class="control">
                  <input v-model="newPhone.number" class="input" type="phone">
                </p>
                <p class="control">
                  <a v-on:click="phoneCreation = !phoneCreation" class="button"> X </a>
                </p>
              </template>
              <template v-else>
                <a v-on:click="phoneCreation = !phoneCreation" class="button">new_phone</a>
              </template>
            </div>
            <div v-if="phoneCreation" class="field">
              <a v-on:click="createPhone">create_phone</a>
            </div>
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
import { userGetRequest, userUpdateRequest, phoneCreateRequest, phoneDeleteRequest } from '@/api/user.js'
export default {
  name: 'EditProfile',
  data () {
    return {
      user: {
        name: ''
      },
      phoneCreation: false,
      newPhone: {
        country_code: 7,
        number: ''
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
        const resp = await userUpdateRequest(this.user.uuid, this.user)
        console.log(resp)
      } catch (e) {
        console.log(e)
      }
    },
    async createPhone (data) {
      try {
        const resp = await phoneCreateRequest({ ...this.newPhone, user_uuid: this.user.uuid })
        console.log(resp.data)
        this.user.phones.push(resp.data)
        this.$store.dispatch('showMsg', {text: 'phone_creation_success', type: 'success'})
      } catch (e) {
        this.$store.dispatch('error', 'phone_creation_error')
        console.log(e)
      }
    },
    async removePhone (phoneUUID) {
      try {
        const resp = await phoneDeleteRequest(this.user.uuid, phoneUUID)
        console.log(resp.data)
        this.$store.dispatch('showMsg', {text: 'phone_remove_success', type: 'success'})
        this.user.phones = this.user.phones.filter(phone => phone.uuid !== phoneUUID)
        this.user.default_phone = resp.data.default_phone
      } catch (e) {
        // TODO: type of failure (service, or constraints)
        this.$store.dispatch('error', 'phone_remove_failure')
        console.log(e)
      }
    },
    changePassword () {
      this.$router.push()
    },
    send () {
      // TODO: validation
      this.update()
    }
  },
  created () {
    console.log('in created')
    this.getUser()
  }
}
</script>

