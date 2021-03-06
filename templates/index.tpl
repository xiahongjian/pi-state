<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>首页</title>

    <link rel="stylesheet" href="static/bootstrap.min.css" />

</head>
<body>
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">Pi State</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
            <li class="active"><a href="#">Home</a></li>
            <!--
            <li><a href="#about">About</a></li>
            <li><a href="#contact">Contact</a></li>
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
              <ul class="dropdown-menu">
                <li><a href="#">Action</a></li>
                <li><a href="#">Another action</a></li>
                <li><a href="#">Something else here</a></li>
                <li role="separator" class="divider"></li>
                <li class="dropdown-header">Nav header</li>
                <li><a href="#">Separated link</a></li>
                <li><a href="#">One more separated link</a></li>
              </ul>
            </li>
            -->
          </ul>
        </div><!--/.nav-collapse -->
      </div>
    </nav>
    <div class="container" style="margin-top: 51px;">
      <div class="jumbotron" style="margin-top: 10px;">
        <h2>基本信息</h2>
        <ul>
          <li>系统: {{ .Data.OS }}</li>
          <li>运行时间: {{ .Data.Uptime }}</li>
          <li>温度: {{ .Data.Temperature }}</li>
          <li>CPU: {{ .Data.CPUName }}</li>
          <li>CPU核心数: {{ .Data.CPUCoreCount }}</li>
          <li>内存总量: {{ .Data.MemoryTotal }}MB</li>
          <li>内存已使用: {{ .Data.MemoryUsed }}MB</li>
          <li>内存已使用占比: {{ .Data.MemoryPercent }}%</li>
          <li>硬盘总量: {{ .Data.DiskTotal }}MB</li>
          <li>硬盘已使用: {{ .Data.DiskFree }}MB</li>
          <li>硬盘已使用占比: {{ .Data.DiskUsedPercent }}%</li>
        </ul>
      </div>
    </div>
<!--
<script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
<script type="text/javascript" src="static/bootstrap.min.js"></script>
-->
</body>
</html>