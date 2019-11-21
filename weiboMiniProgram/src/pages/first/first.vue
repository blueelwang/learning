<template>
  <view class="container">
    <view>
      <text class="text" selectable="true">{{text}}</text>
      <icon type="download"></icon>
      <icon type="search"></icon>

      <icon type="clear"></icon>
      <icon type="cancel"></icon>
    </view>
    <view>
      <progress
        active="false"
        active-mode="forwards"
        :percent="percent"
        border-radius="3"
        strokeWidth="6"
      ></progress>
      <button @click="percent ++">+1 当前进度：{{percent}}%</button>
      <button @click="percent --" size="mini">-1 当前进度：{{percent}}%</button>
    </view>

    <template>
      <checkbox value="cb" checked/>选中
      <checkbox value="cb"/>未选中
      <checkbox value="cb" :color="color" checked/>颜色
      <checkbox value="cb" :disabled="disabled"/>不可用
    </template>

    <view>
      <form>
        <input type="text" maxlength="10" placeholder="文本输入键盘">
        <input type="number" maxlength="10" placeholder="数字输入键盘">
        <input type="digit" maxlength="10" placeholder="带小数点的数字键盘">
        <label class="label1" v-for="item in radioitems" :key="item.name">
            <radio key="country" :value="item.value" >
              <text>{{item.name}}</text>
            </radio>
          </label>
      </form>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      text: "我是text，后面是icon: ",
      percent: 27,
      color: "black",
      checked: true,
      disabled: true,
      radioitems: [
        {name: 'USA', value: '美国'},
        {name: 'CHN', value: '中国', checked: 'true'},
        {name: 'BRA', value: '巴西'},
        {name: 'JPN', value: '日本'},
        {name: 'ENG', value: '英国'},
        {name: 'TUR', value: '法国'},
      ]
    };
  },
  methods: {}
};
</script>