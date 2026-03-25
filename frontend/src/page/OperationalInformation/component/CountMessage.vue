<template>
  <div class="countMessage">
    <messageIcon
        :color="color"
    />
    <div><h2
        :style="{ color: color }"
    >
      {{ count }}
    </h2></div>
    <div><span class="finePrint">{{ title }}</span></div>
  </div>
</template>

<script>
import MessageIcon from "@/components/icons/MessageIcon.vue";

export default {
  name: "countMessage",
  props:{
    color:{type: String, default:"#ffffff"},
    count:{type: String},
    title:{type: String}
  },
  components: {MessageIcon}
}
</script>

<style scoped>
.countMessage{
  display: flex;
  align-items: end;
  gap: 15px;
  svg{
    width: 40px;
    height: 40px;
  }
}
</style>