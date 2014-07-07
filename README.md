cflog2ltsv
==========

Converter from CloudFront access log to LTSV

Usage
-----

```
$ cflog2ltsv < cloudfront.access.log
```

A example of input log.
```
#Version: 1.0
#Fields: date time x-edge-location sc-bytes c-ip cs-method cs(Host) cs-uri-stem sc-status cs(Referer) cs(User-Agent) cs-uri-query cs(Cookie) x-edge-result-type x-edge-request-id x-host-header cs-protocol cs-bytes
2014-03-07	02:15:03	NRT52	31813	192.0.2.1	GET	example.cloudfront.net	/path/to/file	200	-	Mozilla/5.0	foo=bar	-	Hit	bUpx3RRZGdqQ1PCm+JKg6E6vcm1pDMHpONfpmLWltD6Ns/+r10vb	example.cloudfront.net	http	273
```

Output in LTSV format.
```
date:2014-03-07	time:02:15:03	x_edge_location:NRT52	sc_bytes:31813	c_ip:192.0.2.1	cs_method:GET	cs_host:example.cloudfront.net	cs_uri_stem:/path/to/file	sc_status:200	cs_referer:-	cs_user_agent:Mozilla/5.0	cs_uri_query:foo=bar	cs_cookie:-	x_edge_result_type:Hit	x_edge_request_id:bUpx3RRZGdqQ1PCm+JKg6E6vcm1pDMHpONfpmLWltD6Ns/+r10vb	x_host_header:example.cloudfront.net	cs_protocol:http	cs_bytes:273
```

LICENSE
-------

The MIT License (MIT)
