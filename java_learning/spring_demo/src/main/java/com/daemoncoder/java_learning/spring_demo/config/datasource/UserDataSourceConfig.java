package com.daemoncoder.java_learning.spring_demo.config.datasource;

import javax.sql.DataSource;

import com.alibaba.druid.spring.boot.autoconfigure.DruidDataSourceBuilder;

import org.apache.ibatis.session.SqlSessionFactory;
import org.mybatis.spring.SqlSessionFactoryBean;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.jdbc.datasource.DataSourceTransactionManager;
import org.springframework.transaction.PlatformTransactionManager;

@Configuration
@MapperScan(basePackages = {"com.daemoncoder.java_learning.spring_demo.dao"}, sqlSessionFactoryRef = "userSqlSessionFactory")
public class UserDataSourceConfig {
    
    @Value("${spring.user-datasource.url:#{null}}")
    private String dataSourceUrl;

    @Primary
    @Bean(name = "userDatasource")
    @ConfigurationProperties("spring.user-datasource")
    public DataSource mysqlDataSource() {
        return DruidDataSourceBuilder.create().build();
    }

    @Primary
    @Bean(name = "userSqlSessionFactory")
    public SqlSessionFactory mysqlSqlSessionFactory(@Qualifier("userDatasource") DataSource dataSource)
            throws Exception {
        final SqlSessionFactoryBean sessionFactory = new SqlSessionFactoryBean();
        sessionFactory.setDataSource(dataSource);
        sessionFactory.getObject().getConfiguration().setMapUnderscoreToCamelCase(true);
        return sessionFactory.getObject();
    }

    @Bean(name = "userTransactionManager")
    public PlatformTransactionManager prodTransactionManager(@Qualifier("userDatasource") DataSource dataSource) {
        return new DataSourceTransactionManager(dataSource);
    }
}
