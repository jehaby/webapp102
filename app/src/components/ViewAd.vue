<template>
  <div>
    <h5> View Ad </h5>
    <p>name: {{ ad.name }}</p>
    <p>description: {{ ad.description }}</p>
    <p>uuid: {{ $route.params.uuid }}</p>
  </div>
</template>

<script>
  import { getAd } from './../api/ad.js'

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
    async beforeRouteEnter (to, from, next) {
      let resp = {}
      try {
        console.log('in hook: ', to.params.uuid)
        resp = await getAd(to.params.uuid)
        console.log('got response', resp)
        next(vm => vm.setAd(resp.data.ad))
      } catch (e) {
        // TODO: better errors
        console.log('couldnd load data ', e)
      }
    },
    methods: {
      setAd (ad) {
        this.ad = ad
      }
    }
  }
</script>

<style scoped>

</style>
