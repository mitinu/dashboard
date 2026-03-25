<template>
  <div>
    <main-header
        :title="titleHeader"
    />
    <main-body/>
  </div>
</template>

<script>
import mainHeader from "@/components/MainHeader.vue";
import mainBody from "@/components/MainBody.vue";
export default {
  name:"app",
  components:{
    mainHeader,
    mainBody
  },
  data(){
    return{
      titleHeader:"реализация национальных проектов",
    }
  }
}
</script>

<style>
a{
  color: aliceblue;
}
.gradientLine {
  height: 1px;
  background: linear-gradient(90deg, rgba(96, 96, 96, 0.3), rgba(96, 96, 96, 0.6), rgba(96, 96, 96, 0.3));
}
h2{
  font-size: 22pt;
}
h3{
  font-size: 18pt;
}
h4{
  font-size: 16pt;
}
.mt20{
  margin-top: 20px;
}
.mt25{
  margin-top: 20px;
}
*{
  border: none;
  margin: 0;
  padding: 0;
  font-size: 14pt;
}
.finePrint{
  font-size: 12pt;
  font-weight: lighter;
  color: lightsteelblue;
}
a{
  text-decoration: none;
}
body{
  background-color: #0b0b10;
  color: aliceblue;
}
input::placeholder {
  color: #757575;
  opacity: 1;
}
</style>
