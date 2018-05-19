<template>
  <div class="field-body">
    <div class="field is-narrow">
      <div class="control">
        <b-autocomplete
            class="is-narrow"
            :data="values"
            placeholder="..."
            field="name"
            maxlenght="10"
            @input="option => selected = option"
        >
        </b-autocomplete>
        <p v-if="msg" class="help is-danger">
          {{ msg }}
        </p>
      </div>
    </div>
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
      selected: '',
      msg: ''
    }
  },
  props: [
    'values'
  ],
  watch: {
    selected (chosen) {
      this.msg = (chosen !== '' && !this.values.includes(chosen))
        ? (this.msg = 'property.unknown_property: ' + chosen)
        : ''

      this.$emit('input', chosen)
    }
  }

}
</script>

