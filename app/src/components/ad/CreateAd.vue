<template>
  <div>
    <h5> Create Ad </h5>
    <form>
      <p>
        name
        <input required v-model="ad.name" type="text"/>
      </p>
      <p>
        description
        <textarea required v-model="ad.description"/>
      </p>
      <p>
        <input type="button" v-on:click="create" value="create"/>
      </p>
    </form>
  </div>
</template>

<script>
  import { createAd } from '../../api/ad.js'

  export default {
    name: 'CreateAd',
    data () {
      return {
        ad: {
          name: '',
          description: ''
        }
      }
    },
    methods: {
      async create () {
        // TODO: form validation
        let resp = {}
        try {
          console.log('in try component', this.$store.state.auth.jwtToken)
          resp = await createAd({...this.ad}, this.$store.state.auth.jwtToken)
        } catch (e) {
          console.log('ad creation failed', e)
          // TODO: better errors
          return this.$store.dispatch('error', 'Ad creation failed')
        }
        console.log('got resp in craeteAD: ', resp)
        this.$router.push('/ads/' + resp.data.ad.uuid)
      }
    }
  }
</script>

<style scoped>

</style>
