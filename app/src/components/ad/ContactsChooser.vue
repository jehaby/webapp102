<template>

    <div class="field is-expanded">
      <div class="field has-addons">
        <p class="control">
          <a class="button is-static">
            {{ country_code }}
          </a>
        </p>
        <p class="control">
          <input class="input" type="tel" placeholder="Your phone number">
        </p>
      </div>
      <p class="help">Do not enter the first zero</p>
    </div>
    <!-- <div class="control">
        <div class="select">
            <input type="tel">
        </div> -->
    </div>

</template>

<script>
// TODO: get localities from backend
export default {
  name: 'ContactsChooser',
  data () {
    return {
      country_code: '+7',
      data: {},
      selected: ''
    }
  },
  watch: {
    selected (chosen) {
      console.log('in contacts selected', chosen)
      this.$emit('chosen', chosen)
    }
  }
}
</script>
