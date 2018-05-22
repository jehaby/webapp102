<template>
<div>
            <div v-for="ad in ads.edges">
                <a :href="ad.node.uuid" class="href">
                    {{ad.node.name}}
                </a>
                <p>{{ad.node.description}}</p>
                <p>{{(ad.node.price / 100).toString() + ' '+ currency }}</p> 
            </div>

</div>
</template>

<script>
import ADS_FILTER from './../../graphql/AdsFilter.gql'

export default {
  name: 'MyAds',
  components: {},
  data () {
    return {
      ads: {}
    }
  },
  apollo: {
    ads: {
      query: ADS_FILTER,
      variables () {
        return {
          args: this.adsArgsAPI
        }
      },
      update ({ ads }) {
        console.log('from ads update', ads)
        return ads
      },
      error (e) {
        //  TODO: show error
        console.log('error in gkl', e)
      }
    }
  }

}
</script>

