<template>
<div>
    <div class="borders">
        <category-chooser v-on:chosen="categoryChosen">
        </category-chooser>
    </div>

    <div class="columns">
        <div class="column is-one-quarter">

            <div>
                <label class="label">sort by</label>
                <div class="field is-grouped">
                    <div class="control">
                        <div class="select">
                            <select v-model="adsArgs.order.orderBy">
                                <option value="DATE">date</option>
                                <option value="PRICE">price</option>
                                <!-- TODO: weight not gonna needed for services! -->
                                <option value="WEIGHT">weight</option>
                            </select>
                        </div>
                    </div>
                    <div class="control">
                        <div class="select">
                            <select v-model="adsArgs.order.direction">
                                <option value="ASC">asc</option>
                                <option value="DESC">desc</option>
                            </select>
                        </div>
                    </div>  
                </div>
            </div>

            <div>
                <label class="label">price</label>
                <div class="field has-addons has-addons-centered">
                    <p class="control"><a class="button is-static">₽</a></p>
                    <p class="control">
                        <input v-model="adsArgs.price.min" class="input" type="number" placeholder="price_from">
                    </p>
                </div>
                <div class="field has-addons has-addons-centered">
                    <p class="control"><a class="button is-static">₽</a></p>
                    <p class="control">
                        <input v-model="adsArgs.price.max" class="input" type="number" placeholder="price_to">
                    </p>
                </div>
            </div>

            <div>
                <strong>locality</strong>
            </div>

            <!-- SPECIFIC OPTIONS for products -->
            <div>
                <strong>brand</strong>
            </div>

            <div>
                <strong>weight</strong>
                <div class="field has-addons has-addons-centered">
                    <p class="control"><a class="button is-static"></a></p>
                    <p class="control">
                        <input class="input" type="number" placeholder="weight_from">
                    </p>
                </div>
                <div class="field has-addons has-addons-centered">
                    <p class="control"><a class="button is-static"></a></p>
                    <p class="control">
                        <input class="input" type="number" placeholder="weight_to">
                    </p>
                </div>
            </div>

            <div>
                <strong>condition</strong>
                
                <condition-chooser>
                </condition-chooser>
            </div>

            <div class="field is-grouped">
                <div class="control">
                    <button class="button" v-on:click="filter">
                        filter
                    </button>
                </div>
                <div class="control">
                    <button class="button" v-on:click="resetFilters">
                        resetFilters
                    </button>
                </div>
            </div>

            <!-- SPECIFIC OPTIONS for category -->
            <!-- fork (for example) -->
            <div>
                <strong>axle</strong>
            </div>

            <div>
                <strong>brake mounts</strong>
            </div>

            <div>
                <strong></strong>
            </div>

            <div>

            </div>
            
        </div>


        <div class="column is-three-quarters">
            <div v-for="ad in ads.edges">
                <a :href="ad.node.uuid" class="href">
                    {{ad.node.name}}
                </a>
                <p>{{ad.node.description}}</p>
                <p>{{(ad.node.price / 100).toString() + ' '+ currency }}</p> 
            </div>
        </div>
        

    </div>
</div>
</template>

<script>
import CategoryChooser from './CategoryChooser'
import ConditionChooser from './ConditionChooser'
import ADS_FILTER from './../../graphql/AdsFilter.gql'

const DEFAULT_FILTERS = {
  first: 10,
  after: null,

  categoryId: null,
  order: {
    orderBy: 'DATE',
    direction: 'DESC'
  },
  price: {
    currency: 'RUB',
    min: null,
    max: null
  },
  weight: {
    min: null,
    max: null
  }
}

export default {
  name: 'ListAds',
  components: { CategoryChooser, ConditionChooser },
  data () {
    return {
      adsArgs: JSON.parse(JSON.stringify(DEFAULT_FILTERS)),
      adsArgsAPI: JSON.parse(JSON.stringify(DEFAULT_FILTERS)),
      ads: [],
      currency: '₽'
    }
  },
  methods: {
    categoryChosen (categoryId) {
      this.adsArgs.categoryId = categoryId
    //   TODO: get category options, show specific filters
    },
    filter () {
      let res = JSON.parse(JSON.stringify(this.adsArgs))
      res.price.min = res.price.min ? parseInt(res.price.min * 100) : null
      res.price.max = res.price.max ? parseInt(res.price.max * 100) : null
      this.adsArgsAPI = res
    },
    resetFilters () {
      let categoryId = this.adsArgs.categoryId
      this.adsArgs = JSON.parse(JSON.stringify(DEFAULT_FILTERS))
      this.adsArgs.categoryID = categoryId
      this.filter()
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
        console.log('from update', ads)
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

<style>
.borders {
  border: dashed;
}
</style>
