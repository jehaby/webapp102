    <template>
      <div class="column is-half">
        <form>

          <div class="field">
            <label class="label">Name</label>
            <div class="control">
              <input v-model="user.name" class="input" type="text">
            </div>
          </div>

          <div class="field">
            <label class="label">Password</label>
            <div class="control">
              <input v-model="user.password" class="input" type="password">
            </div>
          </div>

          <div class="field is-grouped">
            <div class="control">
              <button v-on:click="login" class="button is-link">Login</button>
            </div>
          </div>

        </form>
      </div>
    </template>

    <script>
      import { loginRequest } from '../../api/auth.js'

      export default {
        name: 'Login',
        data () {
          return {
            user: {
              name: '',
              password: ''
            }
          }
        },
        methods: {
          async login () {
            // TODO: form validation
            let resp = {}
            try {
              resp = await loginRequest({...this.user})
              this.$store.commit('user', resp.user)
              this.$store.commit('jwtToken', resp.token)
              this.$router.push('/profile')
            } catch (e) {
              console.log('foooo', e)
              // TODO: better errors
              return this.$store.dispatch('error', 'Login failed')
            }
          }
        }
      }
    </script>

    <style scoped>

    </style>
