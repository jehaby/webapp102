<template>

<div class="field is-horizontal">
  <div class="field-label">
    <label class="label">{{ name }}</label>
  </div>

  <ad-property-bool v-if="type === 'BOOL'"
    v-model="newValue"
  >
  </ad-property-bool>
  <ad-property-range v-if="type === 'RANGE'">
  </ad-property-range>
  <ad-property-values v-if="type === 'VALUES'"
    :values="values"
    v-model="newValue"
  >
  </ad-property-values>

</div>

</template>

<script>

import AdPropertyBool from './AdPropertyBool'
import AdPropertyRange from './AdPropertyRange'
import AdPropertyValues from './AdPropertyValues'

export default {
  name: 'AdProperty',
  components: {AdPropertyBool, AdPropertyRange, AdPropertyValues},
  data () {
    return {
      newValue: this.value,
      ...this.property
    }
  },
  props: [
    'value',
    'property'
  ],
  watch: {
    newValue (chosen) {
      this.$emit('input', chosen)
    }
  }
}
</script>

