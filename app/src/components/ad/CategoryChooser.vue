<template>
  <div>
    <div v-if="done">
      <p>Категория: {{ currentCategoryName }}</p>
      <button v-on:click="change">Изменить</button>
    </div>
    <ul v-else v-for="ids,level in show" class="category-list">
      <li
        v-for="id in ids"
        v-on:click="choose(id)"
        v-bind:class="{ choosen: currentPath.indexOf(id) !== -1 }"
      >
        {{ categories[id].name }}
      </li>
    </ul>
  </div>
</template>

<script>
  import {getCategories} from './../../api/ad.js'

  export default {
    name: 'CategoryChooser',
    data () {
      return {
        done: false,
        currentPath: [],
        show: [],
        categories: {}
      }
    },
    computed: {
      currentCategoryName () {
        return this.categories[this.currentPath[this.currentPath.length - 1]].name
      }
    },
    async created () { // TODO: probably should move to store, use local storage or something for caching
      try {
        this.categories = (await getCategories()).data
      } catch (e) {
        this.$store.dispatch('error', e) // TODO: better err
      }
      this.show = this.siblings(1) // TODO: implied that category with such id exists and on highest level. It might not be so
    },
    methods: {
      choose (categoryId) {
        categoryId = parseInt(categoryId)
        this.currentPath = this.categories[categoryId].path
        this.currentPath.forEach((catId, level) => {
          this.show[level] = this.siblings(catId)
        })
        this.show[this.currentPath.length] = this.children(categoryId)
        if (this.children(categoryId).length === 0) {
          this.done = true // leaf category chosen
        }
      },
      siblings (id) { // TODO: move maybe
        id = parseInt(id) // TODO: Types!
        if (this.categories[id].path.length === 1) {
          return [...new Set(Object.values(this.categories).map(v => v.path[0]))]
        }
        const parentId = this.categories[id].path[this.categories[id].path.indexOf(id) - 1]
        return this.children(parentId)
      },
      children (id) {
        id = parseInt(id) // TODO: Types!
        const res = []
        Object.values(this.categories).forEach(val => {
          let i = val.path.indexOf(id)
          if ((i !== -1) && (val.path.length === (i + 2))) {
            res.push(val.path[i + 1])
          }
        })
        return res // TODO: sort maybe
      },
      change () {
        this.done = false
      }
    }
  }
</script>

<style scoped>
  .category-list {
    display: inline;
    border: 1px;
  }

  .choosen {
    background: lightslategray;
  }

</style>
