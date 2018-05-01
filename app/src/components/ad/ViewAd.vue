<template>
  <div>
    <h5> View Ad </h5>
    <p>name: {{ ad.name }}</p>
    <p>description: {{ ad.description }}</p>
    <p>uuid: {{ $route.params.uuid }}</p>
  </div>
</template>

<script>
  import VIEW_AD from '../../graphql/AdView.gql'

  export default {
    name: 'ViewAd',
    data () {
      return {
        ad: {
          name: '',
          description: ''
        }
      }
    },
    // maybe it's better to use before route enter
    apollo: {
      ad: {
        query: VIEW_AD,
        variables () {
          return {
            uuid: this.$route.params.uuid
          }
        },
        update ({ ad }) {
          return ad
        }
      }
    }
  }
</script>

<style scoped>

</style>
