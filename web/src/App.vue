<template>
  <v-app>
    <v-main id="scroll__main">
      <v-container fluid>
        <router-view></router-view>
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
