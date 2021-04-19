# gqzcl doc

> 记录


<p v-if="false">Text for Github</p>
<ul>
  <li v-for="i in 3">Item {{ i }}</li>
</ul>



layer.msg('hello'); 



layui.use('layer', function(){
  var layer = layui.layer; 
  layer.msg('hello');
}); 

<script src="layer.js"></script>
<script>
 layer.msg('hello'); 
</script>


<div id="app">
	{{ message }}
</div>

<script type="text/javascript">
var app = new Vue({
	el:'#app',
	data:{
		message:'hello'
	}
});
</script>