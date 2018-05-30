<template>
  <div>
    <!--TODO: forbid to show when logged out-->
    <h3> Hello, {{ user.name }} </h3>
    <p> Your profile </p>
    <p> email {{ user.email }}</p>
  </div>
</template>

<script>
  import {mapState} from 'vuex'

  export default {
    name: 'Profile',
    computed: {
      ...mapState({user: state => state.auth.auth.user})
    }
  }
</script>
