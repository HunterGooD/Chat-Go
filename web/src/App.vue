<template>
  <v-app>
    <v-app-bar app dark>
      <v-toolbar-title>
        <v-row align="center">
          <v-col>
            <v-avatar size="50">
              <img
                alt="user"
                src="https://cdn.pixabay.com/photo/2020/06/24/19/12/cabbage-5337431_1280.jpg"
              />
            </v-avatar>
          </v-col>
          <v-col class="hidden-xs-only"> John Doe </v-col>
        </v-row>
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-export</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main id="scroll__main">
      <v-container fluid>
        <v-container fluid>
          <router-view></router-view>
        </v-container>
      </v-container>
    </v-main>

    <v-footer app v-bind="localAttrs" dark :padless="padless">
      <v-card flat tile width="100%" class="lighten-1 text-center">
        <v-bottom-navigation :value="$router.name" color="red" grow>
          <v-btn to="/" icon>
            <span>Главная</span>

            <v-icon>mdi-home</v-icon>
          </v-btn>

          <v-btn to="about" icon>
            <span>Сообщения</span>

            <v-icon>mdi-email</v-icon>
          </v-btn>

          <v-btn icon>
            <span>Каленьдарь</span>

            <v-icon>mdi-calendar</v-icon>
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
