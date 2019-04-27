<template>
  <div class="Tool">
    <form ref="imageform">
      <input
        ref="imageinput"
        type="file"
        name="img"
        @input="onImageInput"
      >
    </form>
    <ImgViewer
      :src="viewerSrc"
      @dbclick="onSelect"
      ref="viewer"
      :background="srcBackground"
    />
    <a
      ref="downloadbtn"
      :href="newimagesrc"
      target="_blank"
      style="display:none"
    />
    <div class="BottomBar">
      <transition
        name="fade"
        mode="out-in"
      >
        <div
          v-if="state=='select'"
          key="1"
        >
          <button @click="onSelect">Select</button>
          <div style="flex-grow:1"></div>
          <button @click="onInference">
            <span>Inference</span>
          </button>
        </div>
        <div
          v-else-if="state=='loading'"
          key="2"
        >
          <svg
            class="icon"
            viewBox="0 0 1024 1024"
          >
            <path
              d="M71.037275 589.62282 343.25474 771.449571 293.661553 848.937992C354.738377 888.720061 426.716444 912.090533 504.085399 912.090533 673.24879 912.090533 817.384122 802.032783 874.652978 647.070874 894.693344 592.803593 926.172549 579.587703 953.425655 587.367905 917.541154 798.956542 733.862684 960.235198 512.059733 960.235198 291.033308 960.235198 107.832701 800.106399 71.037275 589.62282ZM954.396313 443.142974 680.356997 244.277431 723.842501 171.478038C663.527271 133.726886 593.027591 111.59587 517.361022 111.59587 343.971534 111.59587 196.745027 225.775185 141.641486 385.202123 131.800502 402.718777 103.188473 416.890393 75.382838 412.395497 120.749927 212.917692 298.843376 63.884268 512.059733 63.884268 736.087733 63.884268 921.169924 228.418363 954.396313 443.142974Z"
              p-id="697"
            ></path>
          </svg>
        </div>
        <div
          v-else-if="state=='success'"
          key="3"
        >
          <button @click="onBack">Back</button>
          <label class="compare">
            <input
              type="checkbox"
              v-model="compare"
            >Compare
          </label>
          <div style="flex-grow:1"></div>
          <button @click="onSave">Save</button>
        </div>
        <div
          v-else
          key="4"
        >
          <span>Failed</span>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
import ImgViewer from "./ImgViewer.vue";
import axios from "axios";
export default {
  name: "Tool",
  components: {
    ImgViewer
  },
  data() {
    this._mouse = {};
    return {
      imagesrc: "",
      imagename: "",
      newimagesrc: "",
      state: "select",
      compare: false,
      srcBackground: [
        {
          color: "#c4c4c4",
          bgtext: "Double Click To Select Image",
          lutext: "before"
        },
        {
          color: "#bbbbbb",
          bgtext: "Inferenced Photo",
          lutext: "after"
        }
      ]
    };
  },
  watch: {
    state(state) {
      if (state == "select") {
        window.URL.revokeObjectURL(this.newimagesrc);
        this.newimagesrc = "";
      }
    }
  },
  computed: {
    viewerSrc() {
      switch (this.state) {
        case "success":
          if (this.compare) {
            return [this.imagesrc, this.newimagesrc];
          } else {
            return [null, this.newimagesrc];
          }
        default:
          return [this.imagesrc, null];
      }
    }
  },
  methods: {
    async onImageInput() {
      if (this.imagesrc != "") {
        window.URL.revokeObjectURL(this.imagesrc);
      }
      const files = this.$refs.imageinput.files;
      if (files.length > 0) {
        this.imagesrc = window.URL.createObjectURL(files[0]);
      }
    },
    onSelect() {
      if (this.state == "select") {
        this.$refs.imageinput.click();
      } else if (this.state == "success") {
        this.$refs.viewer.Scalevalue.imgScale = 1;
      }
    },
    onBack() {
      this.state = "select";
    },
    async onInference() {
      this.state = "loading";
      // do ajax
      try {
        const req = await axios.post(
          "/inference",
          new FormData(this.$refs.imageform),
          { responseType: "blob" }
        );
        var data = req.data;
        this.newimagesrc = window.URL.createObjectURL(data);
      } catch (e) {
        console.log(e);
        this.state = "failed";
      }
      this.state = "success";
    },
    onSave() {
      let nameList = this.$refs.imageinput.files[0].name.split(".");
      nameList.splice(-1, 0, "inference");
      this.$refs.downloadbtn.download = nameList.join(".");
      this.$refs.downloadbtn.click();
    }
  }
};
</script>

<style lang="scss">
.Tool {
  flex-grow: 1;
  position: relative;
  display: flex;
  form {
    display: none;
  }
  .BottomBar {
    position: absolute;
    bottom: 0;
    width: 100%;
    height: 50px;
    background: rgba(0, 0, 0, 0.8);
    box-shadow: 0px 0px 13px rgba(0, 0, 0, 0.9);
    color: white;
    > div {
      display: flex;
      align-items: center;
      justify-content: center;
      height: 100%;
      width: 100%;
      button {
        color: white;
        background: transparent;
        border-radius: 3px;
        border: 1px solid white;
        padding: 5px 15px;
        font-weight: lighter;
        font-size: 15px;
        transition: background-color ease-in-out 0.3s;
        margin-left: 10px;
        &:focus {
          outline: none;
        }
        &:hover {
          background: rgba(255, 255, 255, 0.3);
        }
        &:active {
          background: rgba(255, 255, 255, 0.7);
        }
        &:nth-child(1) {
          margin-left: 20px;
        }
        &:last-child {
          margin-right: 20px;
        }
      }
      .icon {
        fill: currentColor;
        height: 30px;
        animation: iconRotate 2s;
        animation-iteration-count: infinite;
      }
      @keyframes iconRotate {
        from {
          transform: rotate(0deg);
        }
        to {
          transform: rotate(360deg);
        }
      }
    }
  }
  .compare {
    margin-left: 10px;
    font-size: 13px;
    display: flex;
    align-items: center;
    > input {
      -webkit-appearance: none;
      vertical-align: middle;
      margin: 0;
      margin-right: 5px;
      background: white;
      height: 12px;
      width: 12px;
      position: relative;
      &:focus {
        outline: none;
      }
      &:checked::after {
        content: "";
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        background: #3a3a3a;
        animation: scaleIn 0.3s;
        height: 50%;
        width: 50%;
      }
      @keyframes scaleIn {
        from {
          height: 0%;
          width: 0%;
        }
        80% {
          height: 60%;
          width: 60%;
        }
        to {
          height: 50%;
          width: 50%;
        }
      }
    }
  }
}
</style>
