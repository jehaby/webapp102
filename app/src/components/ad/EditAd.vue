<template>
  <div class="column is-8">

    <h1 class="title"> Edit Ad </h1>

    <form>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">category</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <p>{{ ad.category.name }}</p>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">name</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <input v-model="ad.name" class="input" type="text" required>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">description</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <textarea v-model="ad.description" class="textarea" type="text" required></textarea>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">price</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <input v-model.number="ad.price" class="input" type="number" required>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">condition</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <div class="select">
                <select v-model="ad.condition" name="" id="">
                  <option value="USED">used</option>
                  <option value="USED_LIKE_NEW">used_like_new</option>
                  <option value="NEW">new</option>
                  <option value="MALFUNCTIONED">malfunctioned</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">weight</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <input v-model.number="ad.weight" class="input" type="number">
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">brand</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
                <p> TODO </p>
              <!-- <brand-chooser v-on:chosen="brandChosen"></brand-chooser> -->
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">contacts</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <contacts-chooser v-on:chosen="contactsChosen"></contacts-chooser>
            </div>
          </div>
        </div>
      </div>
      

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">locality</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <locality-chooser v-on:chosen="localityChosen"></locality-chooser>
            </div>
          </div>
        </div>
      </div>

      <template v-for="property in categoryProperties">
        <ad-property
          :property="property"
          v-model="ad.properties[property.name]"
        >
        </ad-property>
      </template>      

      <div class="field is-horizontal">
        <div class="field-label">
          <!-- Left empty for spacing -->
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <button v-on:click="update" :disabled="! ready" class="button is-primary">
                save
              </button>
            </div>
          </div>
        </div>
      </div>

    </form>
  </div>
</template>

<script>
import AdProperty from './property/AdProperty'
import BrandChooser from './BrandChooser'
import ContactsChooser from './ContactsChooser'
import LocalityChooser from './LocalityChooser'
import { difference } from './../../util/diff.js'
import { isEmpty } from 'lodash'
import AD_UPDATE from './../../graphql/AdUpdate.gql'
import PROPERTIES from './../../graphql/Properties.gql'
import VIEW_AD from '../../graphql/AdView.gql'

export default {
  name: 'EditAd',
  components: { LocalityChooser, BrandChooser, ContactsChooser, AdProperty },
  data () {
    return {
      ad: {
        // TODO: empty ad or fetch before component
        name: '',
        category: {
          name: ''
        }
      },
      originalAd: {},
      categoryProperties: {}
    }
  },
  computed: {
    updateArg () {
      if (isEmpty(this.ad) || isEmpty(this.originalAd)) {
        return {}
      }

      let res = difference(this.ad, this.originalAd)
      if (res.price) {
        res.price = parseInt(res.price * 100)
      }
      if (res.weight) {
        res.weight = parseInt(res.weight)
      }
      if (res.properties) {
        res.properties = JSON.stringify(this.ad.properties)
      }
      return res
    },
    ready: function () {
      return !isEmpty(this.updateArg)
    }
  },
  methods: {
    localityChosen (id) { this.ad.localityId = id.toString() },
    brandChosen (id) { this.ad.brandId = id.toString() },
    contactsChosen (c) {
      this.ad.phoneUUID = c
    },
    async update () {
      try {
        let resp = await this.$apollo.mutate({
          mutation: AD_UPDATE,
          variables: {'uuid': this.ad.uuid, 'input': this.updateArg},
          client: 'auth'
        // update: (store, { data: { adCreate } }) => {
        // }
        })
        console.log('resp ok', resp)
        // TODO: bugfix view shows old values (select before update)
        this.$router.push('/ads/view/' + resp.data.adUpdate.uuid)
      } catch (e) {
        console.log('resp err', e)
        return this.$store.dispatch('error', 'Ad update failed')
      }
    }
  },
  apollo: {
    ad: {
      query: VIEW_AD,
      variables () {
        return {
          uuid: this.$route.params.uuid
        }
      },
      update ({ ad }) {
        console.log('from update: ', ad)

        this.originalAd = {...ad}
        this.originalAd.price = parseInt(this.originalAd.price / 100)
        // TODO: handle json parsing error (lib maybe?)
        this.originalAd.properties = JSON.parse(this.originalAd.properties)

        return JSON.parse(JSON.stringify(this.originalAd))
      }
    },
    categoryProperties: {
      query: PROPERTIES,
      skip () {
        return !this.ad.category.id
      },
      variables () {
        return {
          categoryId: this.ad.category.id.toString()
        }
      },
      update ({ properties }) {
        return properties.reduce(
            (acc, val) => {
              return Object.assign(acc, {
                [val.name]: {...val}
              })
            },
             {})
      },
      error (e) {
        //  TODO: show error
        console.log('error in apollo properties', e)
      }
    }
  }
}
</script>
