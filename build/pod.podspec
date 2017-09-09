Pod::Spec.new do |spec|
  spec.name         = 'Xenio'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'github.com/xenioplatform/go-xenio'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Xenio Client'
  spec.source       = { :git => 'github.com/xenioplatform/go-xenio.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Xenio.framework'

	spec.prepare_command = <<-CMD
    curl https://github.com/xenioplatform/go-xenio/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Xenio.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
