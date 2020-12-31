<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="player" :fields="['name', 'invetory', 'mines', ]" module="mineservice" />
		<sp-type-form type="resource" :fields="['name', 'price', 'selling', 'minPrice', ]" module="mineservice" />
		<sp-type-form type="mine" :fields="['name', 'price', 'owner', 'selling', 'efficiency', 'invetory', 'resources', 'uraniumCost', ]" module="mineservice" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
