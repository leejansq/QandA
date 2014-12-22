<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.0//EN" "http://www.wapforum.org/DTD/xhtml-mobile10.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
{{with .Page}}
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta content="width=device-width,minimum-scale=1,maximum-scale=1,initial-scale=no" name="viewport">
<title>{{.Head}}</title>
<link rel="stylesheet" type="text/css" href="/static/css/down.css">
<link rel="stylesheet" media="only screen and (max-device-width: 320px)" href="/static/css/iphone.css" type="text/css"/>
<link rel="stylesheet" media="only screen and (min-device-width: 321px) and (max-device-width:1080px)" href="/static/css/android.css" type="text/css" />
</head>

<body>
	<div id="layout">
		<div class="top">
			{{.Head}}
			<a class="ret"></a>
			<a class="search"></a>
		</div>
		<div class="listbox">
			{{range .Body}}
			<div class="list">
				<div class="list_title"><span>&nbsp;</span><h1>{{.Key}}.{{.Qustn}}</h1></div>
				{{range .Contn}}
				<div class="list_con">
					<ul>
						<li>
							<p>{{.Page}}</p>
							{{if .Img}}
							<img src="/static/img/{{.Img}}" />
							{{end}}
						</li>
					</ul>
				</div>
				{{end}}
			</div>
			{{end}}
			<!--
			<div class="list">
				<div class="list_title"><span>&nbsp;</span><h1>1、运营商网络</h1></div>
				<div class="list_con">
					<ul>
						<li>
							<p>小蚁摄像机需要一定的上传带宽维持一个稳定的连接，如
果您的小蚁摄像机端没有足够的上传带宽，我们建议您联
系您的互联网服务提供商，是否提供升级服务。</p>
							<img src="/static/img/listimg1.png" />
						</li>
					</ul>
				</div>
			</div>
			<div class="list">
				<div class="list_title"><span>&nbsp;</span><h1>2、信号干扰</h1></div>
				<div class="list_con">
					<ul>
						<li>
							<p>有可能您的小蚁摄像机附近的其它电子设备干扰了，尽量
远离无绳电话，微波炉，还有一些大屏幕电视或者家庭影
院系统，有时候只要稍微移动几十厘米，效果就不一样了
哦。</p>
							<img src="/static/img/listimg2.png" />
						</li>
					</ul>
				</div>
			</div>
			<div class="list">
				<div class="list_title"><span>&nbsp;</span><h1>3、信号测试</h1></div>
				<div class="list_con">
					<ul>
						<li>
							<p>测试WIFI信号强度，把摄像机移近您的路由器，家庭房间使
用的不同材料也会削弱WIFI信号强度，尽量把小蚁摄像机和
路由器放在同一个房间使用哦。</p>
							<img src="/static/img/listimg3.png" />
						</li>
					</ul>
				</div>
			</div>
			-->
		</div>
	</div>
</body>
{{end}}
</html>
