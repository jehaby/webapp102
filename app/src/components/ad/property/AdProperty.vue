<template>

<div class="field is-horizontal">
  <div class="field-label">
    <label class="label">{{ name }}</label>
  </div>
  <div class="field-body">
    <div class="field is-narrow">
      <div class="control">

        <ad-property-bool v-if="type === 'BOOL'"
        >
        </ad-property-bool>
        <ad-property-range v-if="type === 'RANGE'">
        </ad-property-range>
        <ad-property-values v-if="type === 'VALUES'"
          :values="values"
          v-model="value"
        >
        </ad-property-values>

      </div>
    </div>
  </div>
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
      value: '',
      ...this.property
    }
  },
  props: [
    'property'
  ],
  watch: {
    value (chosen) {
      this.$emit('input', chosen)
    }
  }
}
</script>

