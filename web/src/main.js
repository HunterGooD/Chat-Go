import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false

router.beforeEach((to, from, next) => {

  if (to.meta.auth && !router.app.isAuth) {
    next("/signin")
  } else if ((to.path == "/signin" || to.path == "/signup") && router.app.isAuth) {
    next("/")
  } else {
    next()
  }

});

new Vue({
  data() {
    return {
      isAuth: true,
    }
  },
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')