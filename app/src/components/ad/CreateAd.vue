<template>
  <div class="column is-8">
    <h5> Create Ad </h5>
    <form>
      <category-chooser v-on:chosen="categoryChosen"></category-chooser>

      <div class="field">
        <label class="label">Name</label>
        <div class="control">
          <input v-model="ad.name" class="input" type="text">
        </div>
      </div>

      <div class="field">
        <label class="label">Name</label>
        <div class="control">
          <textarea v-model="ad.description" class="input" type="text"></textarea>
        </div>
      </div>

      <div class="field is-grouped">
        <div class="control">
          <button v-on:click="create" class="button is-link">Create</button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
  import { createAd } from '../../api/ad.js'
  import CategoryChooser from './CategoryChooser'

  export default {
    components: {CategoryChooser},
    name: 'CreateAd',
    data () {
      return {
        ad: {
          name: '',
          description: '',
          category_id: 0
        }
      }
    },
    computed: {
      ready () {
        return this.ad.category_id !== 0
      }
    },
    methods: {
      categoryChosen (id) {
        this.ad.category_id = id
        console.log('in create ad', id)
      },
      async create () {
        // TODO: form validation
        let resp = {}
        try {
          console.log('in try component', this.$store.state.auth.jwtToken)
          resp = await createAd({...this.ad}, this.$store.state.auth.jwtToken)
        } catch (e) {
          console.log('ad creation failed', e)
          // TODO: better errors
          return this.$store.dispatch('error', 'Ad creation failed')
        }
        console.log('got resp in craeteAD: ', resp)
        this.$router.push('/ads/' + resp.data.ad.uuid)
      }
    }
  }
</script>

<style scoped>

</style>
