<template>
    <div>
        <b-autocomplete
            class="is-narrow"
            :data="values"
            placeholder="..."
            field="name"
            maxlenght="10"
            @input="option => selected = option"
        >
            
        </b-autocomplete>
    </div>

</template>

<script>
import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.component(Buefy.Autocomplete.name, Buefy.Autocomplete)
Vue.component(Buefy.Field.name, Buefy.Field)

export default {
  name: 'AdPropertyValues',
  data () {
    return {
      name: '',
      selected: ''
    }
  },
  props: [
    'values'
  ],
  watch: {
    selected (chosen) {
      this.$emit('input', chosen)
    }
  }

}
</script>

