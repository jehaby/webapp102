<template>

    <div class="control">
        <b-field>
            <b-autocomplete
                class="is-narrow"
                v-model="name"
                :data="filteredBrands"
                placeholder="..."
                field="name"
                maxlenght="10"
                @select="option => selected = option">
                <!-- <template slot="empty">No results found</template> -->
            </b-autocomplete>
        </b-field>
    </div>

</template>

<script>
import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import BRANDS_ALL from '../../graphql/Brands.gql'

Vue.component(Buefy.Autocomplete.name, Buefy.Autocomplete)
Vue.component(Buefy.Field.name, Buefy.Field)

export default {
  name: 'BrandChooser',
  data () {
    return {
      name: '',
      selected: '',
      brands: []
    }
  },
  computed: {
    filteredBrands () {
      return this.brands.filter((option) => {
        return option.name
                        .toString()
                        .toLowerCase()
                        .indexOf(this.name.toLowerCase()) >= 0
      })
    }
  },
  watch: {
    selected (chosen) {
      this.$emit('chosen', chosen.id)
    }
  },
  apollo: {
    brands: {
      query: BRANDS_ALL,
      update ({ brands }) {
        return brands
      }
    }
  }
}
</script>
