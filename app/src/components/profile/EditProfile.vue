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
            <input v-model="user.name" class="input" type="email">
          </div>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">phone</label>
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
    changePassword () {
      this.$router.push()
    },
    send () {

    }
  }
}
</script>

