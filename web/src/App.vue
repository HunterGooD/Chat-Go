<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" dark temporary right app>
      <template v-slot:prepend>
        <v-list-item two-line>
          <v-list-item-avatar>
            <img src="https://randomuser.me/api/portraits/women/81.jpg" />
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title>
              <span>Jane Smith </span>
            </v-list-item-title>
            <v-list-item-subtitle>Logged In</v-list-item-subtitle>
          </v-list-item-content>
          <v-list-item-action>
            <v-btn icon>
              <v-icon>mdi-logout</v-icon>
            </v-btn>
          </v-list-item-action>
        </v-list-item>
      </template>

      <v-divider></v-divider>

      <v-list dense>
        <v-list-item v-for="item in items" :key="item.title">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main id="scroll__main">
      <v-container fluid>
        <v-container fluid>
          <router-view></router-view>
        </v-container>
      </v-container>
    </v-main>

    <v-footer app v-bind="localAttrs" dark :padless="padless">
      <v-card flat tile width="100%" class="lighten-1 text-center">
        <v-bottom-navigation
          :value="$router.name"
          color="red"
          shift
          v-if="$root.isAuth"
        >
          <v-btn to="/" icon x-small>
            <span>Главная</span>

            <v-icon>mdi-home</v-icon>
          </v-btn>

          <v-btn to="/messages" icon x-small>
            <span>Сообщения</span>

            <v-icon>mdi-email</v-icon>
          </v-btn>

          <v-btn icon x-small>
            <span>Каленьдарь</span>

            <v-icon>mdi-calendar</v-icon>
          </v-btn>

          <v-btn icon @click.stop="drawer = !drawer" x-small>
            <span>настройки</span>

            <v-icon>mdi-dialpad</v-icon>
          </v-btn>
        </v-bottom-navigation>

        <v-divider></v-divider>

        <v-card-text class="white--text">
          {{ new Date().getFullYear() }} — <strong>GoChat</strong>
        </v-card-text>
      </v-card>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: "App",

  components: {},
  data: () => ({
    value: 0,
    padless: true,
    drawer: false,
    group: null,
    variant: "default",
    items: [
      { title: "Home", icon: "mdi-home-city" },
      { title: "My Account", icon: "mdi-account" },
      { title: "Users", icon: "mdi-account-group-outline" },
    ],
  }),
  watch: {
    group() {
      this.drawer = false;
    },
  },
  computed: {
    localAttrs() {
      const attrs = {};
      if (this.variant === "default") {
        attrs.absolute = false;
        attrs.fixed = false;
      } else {
        attrs[this.variant] = true;
      }
      return attrs;
    },
  },
};
</script>
