<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.0//EN" "http://www.wapforum.org/DTD/xhtml-mobile10.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>FAQ后台</title>
<link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>

<body>
<form action="" method="post" enctype="multipart/form-data">
	<table class="tab" width="600" border="0" cellpadding="0" cellspacing="0">
	  <tr>
		<td width="150" class="lef">分类：</td>
		<td><select name="list">
		<option value="">--------请选择分类问题--------</option>
		<option value="1">视频查看</option>
		<option value="2">设备升级</option>
		<option value="3">设备重置</option>
		<option value="4">WIFI强度</option>
		<option value="5">设备连接</option>
		<option value="6">状态灯含义</option>
		</select></td>
	  </tr>
	 <tr>
		<td class="lef">问题序号：</td>
		<td><input name="key" type="text" /></td>
	  </tr>
	  <tr>
		<td class="lef">问题标题：</td>
		<td><input name="qhead" type="text" /></td>
	  </tr>
	  <tr>
		<td colspan="2" class="tab2">
			<table width="100%" border="0" cellpadding="0" cellspacing="0">
			  <tr>
				<td width="150" class="lef">解答1</td>
				<td><input name="qcontent" type="text"/></td>
				<!--<td><textarea name="qcontent" cols="" rows=""></textarea>-->
				</td>
			  </tr>
			  <tr>
				<td width="150" class="lef">图片1</td>
				<td><input name="qimg" type="file" /></td>
				<td><input name="qimgName" type="text" /></td>
			  </tr>
			 
			</table>
		</td>
	  </tr>
				<td>&nbsp;</td>
				<td><input name="" type="submit" value="提 交"/></td>
			  </tr>
	</table>
</form>
</body>
</html>
