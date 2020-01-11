use strict;
use warnings;

hello();

my $avg= average(10,20,30);

print "Average of 10,20 and 30 is equal to $avg .\n";


print "Enter a DNA: ";
my $DNA = <>; 
chomp ($DNA);
$DNA = myfunct($DNA);
print $DNA, "\n";



print "Enter a DNA: ";
my $DNA2 = <>; 
chomp ($DNA2);
myfunc2(\$DNA2);
print $DNA2, "\n";




sub hello {
	
	print "Hello World!\n";

}


sub average {
	my $size = @_; 
	my $total  =0;
	
	
	for (my $i=0; $i<$size;$i++) {
		$total +=  $_[$i];
		#$total = $total+  $_[$i];
	}
	
	my $avg = $total / $size;
	
	return  $avg;
	
}


sub myfunct{
	my ($DNA) = @_;  
	if($DNA =~/atg/){
		$DNA =~s/atg/ATG/g;  
	}
	
	return $DNA ;
}


sub myfunc2  {
my $DNAref = shift;

if ($$DNAref =~ /atg/){  
	$$DNAref =~ s/atg/ATG/g;
	}

}






