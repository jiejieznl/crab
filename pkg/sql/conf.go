package sql

type Config struct {
	Source string `toml:"source" comment:"数据源类型"`
	MYSQL  MYSQL  `toml:"mysql" comment:"mysql配置"`
	PGSQL  PGSQL  `toml:"pgsql" comment:"postgresql配置"`
	SQLITE SQLITE `toml:"sqlite" comment:"sqlite配置"`
}

type MYSQL struct {
	Port         int    `toml:"port" comment:"数据库服务的端口号"`
	Config       string `toml:"config" comment:"数据库连接的额外配置文件路径或JSON字符串"`
	DbName       string `toml:"sql-name" comment:"要连接的数据库名称"`
	Username     string `toml:"username" comment:"用于连接数据库的用户名"`
	Password     string `toml:"password" comment:"对应的用户密码"`
	Addr         string `toml:"addr" comment:"数据库服务器的地址"`
	MaxIdleConns int    `toml:"max-idle-conns" comment:"数据库连接池的最大空闲连接数"`
	MaxOpenConns int    `toml:"max-open-conns" comment:"数据库连接池的最大打开连接数"`
}

type PGSQL struct {
	Port         int    `toml:"port" comment:"数据库服务的端口号"`
	Config       string `toml:"config" comment:"数据库连接的额外配置文件路径或JSON字符串"`
	DbName       string `toml:"sql-name" comment:"要连接的数据库名称"`
	Username     string `toml:"username" comment:"用于连接数据库的用户名"`
	Password     string `toml:"password" comment:"对应的用户密码"`
	Addr         string `toml:"addr" comment:"数据库服务器的地址"`
	MaxIdleConns int    `toml:"max-idle-conns" comment:"数据库连接池的最大空闲连接数"`
	MaxOpenConns int    `toml:"max-open-conns" comment:"数据库连接池的最大打开连接数"`
}

type SQLITE struct {
	DbPath       string `toml:"sql-path" comment:"数据库路径例如./sql.db"`
	MaxIdleConns int    `toml:"max-idle-conns" comment:"数据库连接池的最大空闲连接数"`
	MaxOpenConns int    `toml:"max-open-conns" comment:"数据库连接池的最大打开连接数"`
}
