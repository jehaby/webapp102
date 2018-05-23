<template>
  <div class="field-body">
    <div class="field is-narrow">
      <div class="control">
        <b-autocomplete
            class="is-narrow"
            :data="values"
            placeholder="..."
            field="name"
            open-on-focus
            maxlenght="10"
            v-model="newValue"
        >
        </b-autocomplete>
        <p v-if="msg" class="help is-warning">
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
import { debounce } from 'lodash'

Vue.component(Buefy.Autocomplete.name, Buefy.Autocomplete)
Vue.component(Buefy.Field.name, Buefy.Field)

export default {
  name: 'AdPropertyValues',
  data () {
    return {
      newValue: this.value
    }
  },
  props: [
    'values',
    'value'
  ],
  watch: {
    newValue: debounce(function (chosen) {
      this.$emit('input', chosen)
    }, 300)
  },
  computed: {
    msg: function () {
      return (this.newValue && !this.values.includes(this.newValue))
        ? ('property.unknown_property: ' + this.newValue)
        : ''
    }

  }
}
</script>

