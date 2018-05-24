import { ApolloClient } from 'apollo-client'
import { HttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import VueApollo from 'vue-apollo'
import Vue from 'vue'

Vue.use(VueApollo)

// Create the apollo client
const defaultClient = new ApolloClient({
  // TODO: read docs
  link: new HttpLink({
    // You should use an absolute URL here
    // TODO: get from config
    uri: 'http://localhost:8899/query'
  }),
  cache: new InMemoryCache(),
  connectToDevTools: true
})

const authClient = new ApolloClient({
  // TODO: read docs
  link: new HttpLink({
    // You should use an absolute URL here
    // TODO: get from config
    uri: 'http://localhost:8899/query',
    credentials: 'include'
  }),
  cache: new InMemoryCache(),
  connectToDevTools: true
})

const apolloProvider = new VueApollo({
  clients: {
    auth: authClient
  },
  defaultClient: defaultClient
})

export default apolloProvider
