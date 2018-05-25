<template>
<div class="navbar-menu">
    <div class="navbar-end">
        <template v-if="loggedIn">
            <router-link class="navbar-item" to="/ads/my">my ads</router-link>          
            <router-link class="navbar-item" to="/profile">profile</router-link>
            <router-link class="navbar-item" to="/ads/create">create ad</router-link>
            <button class="button navbar-item" v-on:click="logout">logout</button>
        </template>
        <template v-else>
            <router-link class="navbar-item" to="/login">login</router-link>
            <router-link class="navbar-item" to="/register">register</router-link>
        </template>
    </div>
</div>
</template>

<script>

export default {
  name: 'NavMenu',
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    }
  },
  methods: {
    async logout () {
      await this.$store.dispatch('logout')
      this.$forceUpdate()
      this.$router.push('/')
    }
  }

}
</script>

