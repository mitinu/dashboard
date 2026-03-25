<template>
  <div class="mainBody">
    <div class="content">
      <div class="container">
        <button-return
          class="buttonReturn"
          v-if="visibilityButtonReturn"
          @click="returnPage"
        />
        <div>
          <h2>{{title}}</h2>
        </div>
      </div>
    </div>
    <router-view
        @setTitleMain="setTitle"
        @setVisibilityButtonReturnMain="setVisibilityButtonReturn"
    />
  </div>
</template>

<script>
import buttonReturn from "@/components/UI/ButtonReturn.vue";
export default {
  name: "mainBody",
  components:{
    buttonReturn
  },
  data() {
    return{
      title: "",
      visibilityButtonReturn:false
    }
  },
  methods:{
    setTitle(title){
      this.title = title
    },
    setVisibilityButtonReturn(status){
      this.visibilityButtonReturn = status
    },
    returnPage(){
      this.$router.back()
    }
  }
}
</script>

<style scoped>
.mainBody{
  height: calc(100vh - 90px);
  .content{
    height: 75px;
    display: flex;
    justify-content: space-between;
    box-sizing: border-box;
    padding: 15px;
    .container{
      display: flex;
      justify-content: space-between;
      gap: 20px;
      align-items: center;
    }
  }
}
</style>