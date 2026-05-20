<script>
export default {
  name: "mixinModalWindow",
  data(){
    return{
      visibilityModalWindow: false,
      headerModalWindow: ""
    }
  },
  methods:{

    openModalWindow(){
      this.visibilityModalWindow = true
    },
    closeModalWindow(){
      this.visibilityModalWindow = false
    },
    setHeaderModalWindow(header){
      this.headerModalWindow=header
    }
  }
}
</script>
